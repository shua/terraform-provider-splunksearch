
resource "splunksearch" "SHUA_test_search" {
	name = "SHUA_test_search"
	search = "does/this/work?"
	description = "testing this"
	disabled = true
}
