package main

import "github.com/dex4er/go-tap"

func main() {
    tap.Ok(false, "Ok")
    tap.Is(123, "Aaa", "Is")
    tap.DoneTesting()
}
