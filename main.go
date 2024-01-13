package main

import (
	"os"
)

// the main function receives a  command ,and it's command line arguments, and run the command with the arguments
func main() {
	// the first argument is the command
	//cd to current path: /Users/yang/Evolving
	os.Chdir("/Users/yang/Evolving")

	// the first argument is the command
	projectName := "完善一个叫MMUEvo的代理，能够根据边际效用最大化（MMU）原则，演进目标系统"
	Description := "你将要创建一个系统。这个系统将考虑目标系统的意图，按最大化边际效用的原则，提出一个（business specification）来改进目标系统。然后你需要将业务规格转化为技术需求。最后将技术需求转化为技术规范. 这个技术规范必须是意图完备的,而且必须最大程度保留必要的上下文。以便后续进一步完成自动编码，自动调试。然后重头开始下一轮自动迭代。"
	MMUToTasks(projectName, Description)

	//part 2 并最终转化为目标任务。然后你需要给这个任务开发一个自动完成模块，自动评估模块。直到最终完成这个改进。你可用的工具包括全部的linux 命令。golang 开发环境。一个你可以用来提问的人类专家组。以及llm agent。请确保只有唯一的business specification。
	//run the command

}
