package main

import "github.com/dex4er/go-tap"

func main() {
    tap.Ok(true, "Ok")
    tap.Is("Aaa", "Aaa", "Is")
    tap.Is(123, 123, "Is")
	tap.DoneTesting()
}
