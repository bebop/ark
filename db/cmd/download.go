package cmd

import (
	"github.com/spf13/cobra"
)

/******************************************************************************
ark needs an easy and reproducible way to grab all of the data it uses.
Some people would try to download everything manually with wget or curl but
that's bonkers and I'm not maintaining that. Instead I've made a multi-threaded webscraper.

To be honest most of you reading this will probably just want to use the build
command which will download AND insert all of the data into ark for you
without any intermediary files but I'm leaving this here for the time being until
the build command is implemented.

Download will download and store the latest versions of the following:

	- Rhea RDF file
	- Rhea to Uniprot Sprot mapping file
	- Rhea to Uniprot Trembl mapping file
	- CHEMBL sqlite file
	- Uniprot Sprot XML file
	- Uniprot Trembl XML file
	- All of Genbank's files

Given that these urls are likely not to change I doubt there will be many if
any stability issues from upstream data sources.

TODO: Create flags so that each data source can be downloaded individually.
TODO: Create way to track, report, and resume progress of downloads.

TTFN,
Tim
******************************************************************************/

var downloadCmd = &cobra.Command{
	Use:   "download",
	Short: "Download data for standard deploy build. Run at your own risk.",
	Long:  "Download literally downloads all the base data needed to build a standard ark deployment the amount of data is dummy high to casually test on your personal machine. Run at your own risk.",
	Run: func(cmd *cobra.Command, args []string) {
		// download.ark()
	},
}
