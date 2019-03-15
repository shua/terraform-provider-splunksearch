# building

	go build && terraform init

# debugging

to debug start a proxy eg

	mitmproxy -p 8100

and then run the terraform plan/apply with http_proxy settings

	http_proxy=http://localhost:8100 https_proxy=http://localhost:8100 terraform plan


# config syntax

it will be painful to keep search strings in tf files
we could split the search out to another file (eg "mysearch.splunk"),

	src/
	+ main.tf
	\ mysearch.splunk

	# main.tf
	resource "splunksearch" "mysearch" {
		name = "mysearch"
		search = "${file(mysearch.splunk)}"
	}

I'm not _so_ hot on splitting up the config, just because there's now three
things to change to keep the naming consistent if the name `mysearch` changes
to `myoldsearch` then the resource name is changed, the value `name =` is changed,
and the name of the other file needs to change.

If we wanted to keep all the config in the same file, we could try yml and
loading with https://github.com/Ashald/terraform-provider-yaml .
But then we'd have more renaming problems. meh

