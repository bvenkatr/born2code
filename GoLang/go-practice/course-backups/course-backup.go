package course_backups

import (
	"flag"
	"fmt"
)

func CourseBackups() {
	backupPath := flag.String("course_backup_path", "/tmp/", "path where course will be dumped to")
	flag.Parse()

	fmt.Println("word:", *backupPath)


	fmt.Println("tail:", flag.Args())
}
