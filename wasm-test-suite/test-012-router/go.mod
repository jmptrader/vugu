module github.com/vugu/vugu/wasm-test-suite/test

go 1.14

replace github.com/vugu/vugu => ../..

//replace github.com/vugu/vgrouter => ../../../vgrouter

require (
	github.com/vugu/vgrouter v0.0.0-20200329225024-3b01bdbe25fa
	github.com/vugu/vjson v0.0.0-20191111004939-722507e863cb
	github.com/vugu/vugu v0.1.0
)
