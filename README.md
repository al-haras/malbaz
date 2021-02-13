# Malbaz

A simple CLI tool written in Golang to help quickly download the newest samples from MalwareBazaar.

### Usage

There are a few commands that can be used to specify what is needed. Each command will require an argument. Due to MalwareBazaar's API, you can get specify "time" or "100" when querying `get_recent`. `time` will be samples uploaded in the past hour. `100` will be the last 100 samples uploaded.

- list - Used to list samples for specified range
- download - Used to download samples for specified range
- daily - Used to download yesterday's malware samples as a batch.

#### Example usage:

`malbaz list 100` - lists the last 100 samples uploaded to MalwareBazaar.

`malbaz daily` - downloads the latest collection of daily samples.

`malbaz download time` - downloads all samples uploaded in the last hour.

### Running the files

To run the files you will need to first unzip. The password is "infected". If not, it is typically listed in the tags.

### Disclaimer

This tool was designed with the intent of obtaining fresh malware samples quickly (and without paying for a feed) for evaluating antivirus solutions. Please execise caution when using this tool. The samples are presumably real and will infect your device if you unzip the files and run them. Do not use this on any unauthorized devices or networks.
