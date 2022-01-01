package snowflake

import (
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/sony/sonyflake"
	"os"
	"time"
)

func Snowflake() {
	n, err := snowflake.NewNode(1)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	for i := 0; i < 3; i++ {
		id := n.Generate()
		fmt.Println(
			"id: ", id,
			"node: ", id.Node(),
			"step: ", id.Step(),
			"time: ", id.Time(),
		)
	}
}

func Sonyflake() {
	t, _ := time.Parse("2006-01-02", "2018-01-01")
	settings := sonyflake.Settings{
		StartTime:      t,
		MachineID:      getMachineID,
		CheckMachineID: checkMachineID,
	}

	sf := sonyflake.NewSonyflake(settings)
	id, err := sf.NextID()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(id)
}

func getMachineID() (uint16, error) {
	var machineID uint16
	var err error
	machineID = readMachineIDFromLocalFile()
	if machineID == 0 {
		machineID, err = generateMachineID()
		if err != nil {
			return 0, err
		}
	}

	return machineID, nil
}

func checkMachineID(machineID uint16) bool {
	saddResult, err := saddMachineIDToRedisSet()
	if err != nil || saddResult == 0 {
		return true
	}

	err = saveMachineIDToLocalFile(machineID)
	if err != nil {
		return true
	}

	return false
}

func readMachineIDFromLocalFile() uint16 {
	return 0
}

func generateMachineID() (uint16, error) {
	return 0, nil
}

func saddMachineIDToRedisSet() (uint16, error) {
	return 0, nil
}

func saveMachineIDToLocalFile(id uint16) error {
	return nil
}
