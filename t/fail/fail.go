package main

import "test/tap"

func main() {
    tap.Ok(false, "Ok")
    tap.Is(123, "Aaa", "Is")
    tap.DoneTesting()
}
