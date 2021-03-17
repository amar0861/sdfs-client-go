package main

import (
	context "context"
	"flag"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/olekukonko/tablewriter"
)

//ConfigCmd Configure Volume functions for sdfscli
func ConfigCmd(ctx context.Context, flagSet *flag.FlagSet) {
	cleanstore := flagSet.Bool("cleanup", false, "Returns Volume Info")
	shutdown := flagSet.Bool("shutdown", false, "Shuts down the volume")
	password := flagSet.String("password", "A Password", "Sets The Password for this volume")
	scv := flagSet.Bool("sync-with-cloud", false, "Syncs the Volume with the metadata stored"+
		" in the cloud to make sure both sides are consistent.")
	vcv := flagSet.Int64("sync-from-cloud", 0, "Syncs the Volume with an existing cloud volume")
	bsize := flagSet.String("dse-cache", ".", "Sets the local cache size for the Dedupe Storage Engine")
	lsize := flagSet.String("volume-size", ".", "Sets the local cache size for the Dedupe Storage Engine")
	rspeed := flagSet.String("read-speed", "-1", "Sets the max read speed from storage for blocks in KB/s")
	wspeed := flagSet.String("write-speed", "-1", "Sets the max write speed from storage for blocks in KB/s")
	kage := flagSet.Int64("max-key-age", -1, "Sets the maximum age of a deduplication entry that can be referenced. If set to -1 the age is infinite.")
	vinfo := flagSet.Bool("volume", false, "Returns Volume Info")
	dinfo := flagSet.Bool("dse", false, "Returns Dedupe Storage Info")
	sinfo := flagSet.Bool("system", false, "System Info")
	gcsecd := flagSet.Bool("gc-schedule", false, "Returns The Garbage Collection Schedule")
	ccv := flagSet.Bool("connected-volumes", false, "Returns A list of volumes that are using the same storage")
	levents := flagSet.Bool("events-list", false, "List Events")
	connection := ParseAndConnect(flagSet)

	if *cleanstore == true {

		evt, err := connection.CleanStore(ctx, true, true)
		if err != nil {
			fmt.Printf("Unable to run garbage collection error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Garbage Collection Finished %s \n", evt.ShortMsg)
		return
	}
	if *shutdown == true {

		connection.ShutdownVolume(ctx)
		fmt.Println("Shutting Down Volume")
		return

	}
	if *scv == true {

		evt, err := connection.SyncCloudVolume(ctx, true)
		if err != nil {
			fmt.Printf("Unable to sync cloud volume: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("Syncing with Cloud Completed %s \n", evt.ShortMsg)
		return

	}
	if IsFlagPassed("dse-cache", flagSet) {
		size, err := GetSize(*bsize)
		if err != nil {
			fmt.Printf("Unable to set dse-cache size for: %v error: %v\n", *bsize, err)
			os.Exit(1)
		}
		evt, err := connection.SetCacheSize(ctx, size, true)
		if err != nil {
			fmt.Printf("Unable to set dse-cache size for: %v error: %v\n", *bsize, err)
			os.Exit(1)
		}
		fmt.Printf("%v \n", evt.ShortMsg)

	}
	if IsFlagPassed("volume-size", flagSet) {
		size, err := GetSize(*lsize)
		if err != nil {
			fmt.Printf("Unable to set dse-cache size for: %v error: %v\n", *lsize, err)
			os.Exit(1)
		}
		err = connection.SetVolumeCapacity(ctx, size)
		if err != nil {
			fmt.Printf("Unable to set dse-cache size for: %v error: %v\n", *lsize, err)
			os.Exit(1)
		}
		fmt.Printf("Volume Size Set to : %s \n", FormatSize(size))
	}
	if IsFlagPassed("max-key-age", flagSet) {
		err := connection.SetMaxAge(ctx, *kage)
		if err != nil {
			fmt.Printf("Unable to set max age to : %v ms, error: %v\n", *kage, err)
			os.Exit(1)
		}
		fmt.Printf("Set Max Age To : %v ms\n", *kage)
	}
	if IsFlagPassed("read-speed", flagSet) {
		size, err := strconv.Atoi(*rspeed)
		if err != nil {
			fmt.Printf("Unable to set read speed to: %v error: %v\n", *rspeed, err)
			os.Exit(1)
		}
		err = connection.SetReadSpeed(ctx, int32(size))
		if err != nil {
			fmt.Printf("Unable to set read speed to: %v error: %v\n", *rspeed, err)
			os.Exit(1)
		}
		fmt.Printf("Read Speed Set to : %s \n", FormatSize(int64(size)))
	}

	if IsFlagPassed("write-speed", flagSet) {
		size, err := strconv.Atoi(*wspeed)
		if err != nil {
			fmt.Printf("Unable to set write speed to: %v error: %v\n", *wspeed, err)
			os.Exit(1)
		}
		err = connection.SetWriteSpeed(ctx, int32(size))
		if err != nil {
			fmt.Printf("Unable to set write speed to: %v error: %v\n", *wspeed, err)
			os.Exit(1)
		}
		fmt.Printf("Write Speed Set to : %s \n", FormatSize(int64(size)))
	}
	if IsFlagPassed("sync-from-cloud", flagSet) {
		evt, err := connection.SyncFromCloudVolume(ctx, *vcv, true)
		if err != nil {
			fmt.Printf("Unable to sync cloud volume from %d error: %v\n", *vcv, err)
			os.Exit(1)
		}
		fmt.Printf("Syncing with Cloud Completed %s \n", evt.ShortMsg)
		return
	}
	if IsFlagPassed("password", flagSet) {
		err := connection.SetPassword(ctx, *password)
		if err != nil {
			fmt.Printf("Unable to set password error: %v\n", err)
			os.Exit(1)
		}
		fmt.Println("Password Set")
		return
	}
	if *gcsecd == true {
		cvo, err := connection.GetGCSchedule(ctx)
		if err != nil {
			fmt.Printf("Unable to get GC schedule error: %v\n", err)
			os.Exit(1)
		}
		fmt.Printf("GC Schedule is: %s\n", cvo)
		return
	}
	if *ccv == true {
		cvo, err := connection.GetConnectedVolumes(ctx)
		if err != nil {
			fmt.Printf("Unable to list connected volumes error: %v\n", err)
			os.Exit(1)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"ID", "Hostname", "Version", "Last Checkin", "Size", "Compressed Size", "Local"})
		for _, v := range cvo {
			lastUpdate := ""
			if v.LastUpdate > 0 {
				lastUpdate = time.Unix(0, v.LastUpdate*int64(time.Millisecond)).String()
			}
			table.Append([]string{fmt.Sprintf("%d", v.Id),
				v.Hostname,
				v.SdfsVersion,
				lastUpdate,
				FormatSize(v.Size),
				FormatSize(v.CompressedSize),
				fmt.Sprintf("%t", v.Local),
			})

		}
		table.Render()
		return
	}
	if *sinfo == true {
		svo, err := connection.SystemInfo(ctx)
		if err != nil {
			fmt.Printf("Unable to get system info error: %v\n", err)
			os.Exit(1)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"System Info", ""})
		table.Append([]string{"Active Threads", fmt.Sprintf("%d", svo.ActiveThreads)})
		table.Append([]string{"Blocks Stored", fmt.Sprintf("%d", svo.BlocksStored)})
		table.Append([]string{"CPU Cores", fmt.Sprintf("%f", svo.CpuCores)})
		table.Append([]string{"Free Memory", FormatSize(int64(svo.FreeMemory))})
		table.Append([]string{"Free Space", FormatSize(svo.FreeSpace)})
		table.Append([]string{"Max Blocks", FormatSize(int64(svo.MaxBlocksStored))})
		table.Append([]string{"Sdfs CPU Load", fmt.Sprintf("%.2f%%", svo.SdfsCpuLoad)})
		table.Append([]string{"Total CPU Load", fmt.Sprintf("%.2f%%", svo.TotalCpuLoad)})
		table.Append([]string{"Total Memory", fmt.Sprintf("%f", svo.TotalMemory)})
		table.Append([]string{"Total Storage", FormatSize(svo.TotalSpace)})
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.Render()
		return
	}
	if *levents == true {
		elist, err := connection.ListEvents(ctx)
		if err != nil {
			fmt.Printf("Unable to list events error: %v\n", err)
			os.Exit(1)
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Event ID", "Parent ID", "Level", "Type", "Start Time", "End Time", "Current Count", "Max Count", "Message"})
		for _, v := range elist {
			endTime := ""
			if v.EndTime > 0 {
				endTime = time.Unix(0, v.EndTime*int64(time.Millisecond)).String()
			}
			table.Append([]string{v.Uuid,
				v.ParentUuid,
				v.Level,
				v.Type,
				time.Unix(0, v.StartTime*int64(time.Millisecond)).String(),
				endTime,
				fmt.Sprintf("%d", v.CurrentCount),
				fmt.Sprintf("%d", v.MaxCount),
				v.ShortMsg,
			})

		}
		table.Render()
		return

	}
	if *vinfo == true {
		volumeInfo, err := connection.GetVolumeInfo(ctx)
		if err != nil {
			fmt.Printf("Unable to get volume info error: %v\n", err)
			os.Exit(1)
		}
		data := [][]string{
			{"ID", strconv.FormatInt(volumeInfo.SerialNumber, 10)},
			{"Name", volumeInfo.Name},
			{"Capacity Formatted", FormatSize(volumeInfo.Capactity)},
			{"Capacity Bytes", strconv.FormatInt(volumeInfo.Capactity, 10)},
			{"Used Formatted", FormatSize(volumeInfo.CurrentSize)},
			{"Used Bytes", strconv.FormatInt(volumeInfo.CurrentSize, 10)},
			{"Compressed Formatted", FormatSize(volumeInfo.DseCompSize)},
			{"Compressed Bytes", strconv.FormatInt(volumeInfo.DseCompSize, 10)},
			{"Duplicate Formatted", FormatSize(volumeInfo.DuplicateBytes)},
			{"Duplicate Bytes", strconv.FormatInt(volumeInfo.DuplicateBytes, 10)},
			{"Max Percentage Full", fmt.Sprintf("%.2f%%", volumeInfo.MaxPercentageFull*100)},
			{"Files", strconv.FormatInt(volumeInfo.Files, 10)},
			{"Offline", strconv.FormatBool(volumeInfo.Offline)},
			{"Allow External Links", strconv.FormatBool(volumeInfo.AllowExternalLinks)},
			{"Compress Metadata", strconv.FormatBool(volumeInfo.CompressedMetaData)},
			{"Use Perf Mon", strconv.FormatBool(volumeInfo.UsePerfMon)},
			{"Perfmon File", volumeInfo.PerfMonFile},
			{"Volume Path", volumeInfo.Path},
			{"Read Bytes", fmt.Sprintf("%f", volumeInfo.ReadBytes)},
			{"Read OPS", fmt.Sprintf("%f", volumeInfo.ReadOps)},
			{"Read Timeout (s)", fmt.Sprintf("%d", volumeInfo.ReadTimeoutSeconds)},
			{"Read Errors", strconv.FormatInt(volumeInfo.ReadErrors, 10)},
			{"Write Bytes", fmt.Sprintf("%d", volumeInfo.WriteBytes)},
			{"Write OPS", fmt.Sprintf("%f", volumeInfo.WriteOps)},
			{"Write Timeout (s)", fmt.Sprintf("%d", volumeInfo.WriteTimeoutSeconds)},
			{"Write Errors", strconv.FormatInt(volumeInfo.WriteErrors, 10)},
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Volume Info", ""})
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetAutoMergeCells(true)
		for _, v := range data {
			table.Append(v)
		}
		table.Render()
		return
	}
	if *dinfo == true {
		dInfo, err := connection.DSEInfo(ctx)
		if err != nil {
			fmt.Printf("Unable to get volume info error: %v\n", err)
			os.Exit(1)
		}
		data := [][]string{
			{"Cache Size Formatted", FormatSize(dInfo.CacheSize)},
			{"Cache Size Bytes", strconv.FormatInt(dInfo.CacheSize, 10)},
			{"Max Cache Size Formatted", FormatSize(dInfo.MaxCacheSize)},
			{"Max Cache Size Bytes", strconv.FormatInt(dInfo.MaxCacheSize, 10)},
			{"Current Size Formatted", FormatSize(dInfo.CurrentSize)},
			{"Current Size Bytes", strconv.FormatInt(dInfo.CurrentSize, 10)},
			{"Compressed Size Formatted", FormatSize(dInfo.CompressedSize)},
			{"Compressed Size Bytes", strconv.FormatInt(dInfo.CompressedSize, 10)},
			{"HashTable Entries", strconv.FormatInt(dInfo.Entries, 10)},
			{"Max Page Size", strconv.FormatInt(dInfo.PageSize, 10)},
			{"Storage Driver", dInfo.StorageType},
			{"Cloud Access Key", dInfo.CloudAccessKey},
			{"Cloud Secret Key", dInfo.CloudSecretKey},
			{"Encryption IV", dInfo.EncryptionIV},
			{"Encryption Key", dInfo.EncryptionKey},
			{"Listen Host", dInfo.ListenHost},
			{"Listen Port", fmt.Sprintf("%d", dInfo.ListenPort)},
			{"Listen Encrypted", fmt.Sprintf("%t", dInfo.ListenEncrypted)},
			{"Read Speed", fmt.Sprintf("%d KB/s", dInfo.ReadSpeed)},
			{"Write Speed", fmt.Sprintf("%d KB/s", dInfo.WriteSpeed)},
			{"Max Key Age", fmt.Sprintf("%d ms", dInfo.MaxAge)},
		}
		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"DSE Info", ""})
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetAutoMergeCells(true)
		for _, v := range data {
			table.Append(v)
		}
		table.Render()
		return
	}

}
