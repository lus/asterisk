package static

var (
	// Mode defines whether or not the bot runs in production mode
	Mode = "dev"

	// Version defines the version string of the current instance
	Version = "dev"

	// MongoDatabase defines the MongoDB database name
	MongoDatabase = "asterisk"

	// HastebinURL defines the URL of the hastebin instance to use
	HastebinURL = "https://hasteb.in/"

	// RTeXURL defines the URL of the rTeX API to use
	RTeXURL = "https://rtex.probablyaweb.site/api/v2"

	// MathJSURL defines the URL of the MathJS API to use
	MathJSURL = "http://api.mathjs.org/v4/"

	// LaTeXTemplate defines the template to use for LaTeX expression rendering
	LaTeXTemplate = "\\documentclass[border=2pt]{standalone} \\usepackage[utf8]{inputenc} \\usepackage{amsmath} \\usepackage{xcolor} \\begin{document} \\color{white} #CONTENT# \\end{document}"

	// IntervalRegexString defines the RegEx an interval has to match
	IntervalRegexString = "(\\[|\\()(\\d+),\\s*(\\d+)(\\]|\\))"
)
