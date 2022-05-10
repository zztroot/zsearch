package options

type CmdOptions struct {
	Value   string `short:"d" long:"data" description:"Value to be searched [required]"`
	Path    string `short:"p" long:"path" description:"Specifies the search path, which defaults to the current path"`
	File    bool   `short:"f" long:"file" description:"Search file name"`
	Content bool   `short:"c" long:"content" description:"Search in file content"`
	Version bool   `short:"v" long:"version" description:"Print version info"`
	About   bool   `short:"a" long:"about" description:"View author information"`
	Ext     bool   `short:"e" long:"ext" description:"Supported file formats for reading"`
}
