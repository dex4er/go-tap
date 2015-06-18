package tap

import (
	"fmt"
	"io"
	"os"
    "runtime"
    "strings"
)

var (
	done    bool = false
	guarded bool = false
	failed  int  = 0
	test    int  = 0
	tests   int  = 0
)

func Diag(message string) {
	var buf string = "# " + strings.Replace(message, "\n", "\n# ", -1)
    _print(os.Stderr, buf)
}

func DoneTesting() {
    if done {
        Fail("DoneTesting() was already called")
        return
	}

    if tests == 0 {
        tests = test
        _print(os.Stdout, fmt.Sprintf("1..%d", tests))
    }

    if tests == 0 {
        Diag("No tests run!")
        failed = 255
    }

    if failed > 0 {
        os.Exit(failed)
    }

    done = true
}

func Fail(test_name string) {
    _ok(false, test_name)
}

func Is(got interface{}, expected interface{}, test_name string) {
    if len(test_name) == 0 {
		test_name = fmt.Sprintf("A %T object", got)
	}
    test_name = fmt.Sprintf("%s is %#v", test_name, expected)
    check := got == expected
	_ok(check, test_name)
    if ! check {
        Diag(fmt.Sprintf("         got: %#v\n    expected: %#v\n", got, expected))
    }
}

func Ok(check bool, test_name string) {
    _ok(check, test_name)    
}

func _guard() {
    if test > 0 && tests == 0 && ! done {
        fmt.Fprintln(os.Stderr, "# Tests were run but no plan was declared and DoneTesting() was not seen.") // stderr
	}

    if ! done {
        if failed > 0 {
           failed = 255 - test
           os.Exit(failed)
		}
    }
}

func _ok(check bool, test_name string) {
    test += 1

    message := ""

    if ! check {
        message = "not "
        failed += 1
    }

    message += fmt.Sprintf("ok %d", test)

    if len(test_name) > 0 {
        // if isinstance(test_name, int) or '{0}'.format(test_name).isdigit():
        //    diag("    You named your test '{0}'.  You shouldn't use numbers for your test names.\n    Very confusing.".format(test_name))        
        message += " - " + test_name
    }

    _print(os.Stdout, message)

    if ! check {
        if len(test_name) > 0 {
            Diag("  Failed test " + test_name)
        } else {
            Diag("  Failed test")
		}
        buf := make([]byte, 1024)
        runtime.Stack(buf, false)
        Diag(string(buf))
	}
}

func _print(w io.Writer, message string) {
    fmt.Fprintln(w, message)
}
