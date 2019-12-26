package main

func main() {
	a := App{}
	a.Initialize(
		"dbu",
		"dbu",
		"dbu")

	a.Run(":9001")
}
