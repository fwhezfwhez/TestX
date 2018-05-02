package _case

import "log"

func init() {
	log.SetPrefix("test ")
	log.SetFlags(log.LstdFlags|log.Llongfile)
}
