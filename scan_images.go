package main

import(
    "flag"
    "os"
    "fmt"
    "io/ioutil"
    "strings"
    "image/png"
    "path"
    )
func main() {
  if len(os.Args)==1{
    fmt.Fprintf(os.Stderr,"wrong number of arguments, expected at least a file or folder path\n")
      os.Exit(1)
  }
onlydata := flag.Bool("t",false,"only display data with tabs")
            flag.Parse()
            scanDir := os.Args[len(os.Args)-1]
            stats, err := os.Stat(scanDir)
            if err != nil{
              fmt.Fprintf(os.Stderr,"Error, the path %s is not accessible\n",scanDir)
                os.Exit(2)           
            }
          if(!*onlydata && stats.IsDir()){
            fmt.Printf("Listing image files in %s\n",scanDir)
          }

          if(!*onlydata && !stats.IsDir()){
            fmt.Printf("examining target %s...\n",scanDir)
          }
          if stats.IsDir() {
            files, _ := ioutil.ReadDir(scanDir)
              for _, f := range files {
                ParseFile(path.Join(scanDir,f.Name()),onlydata)
              }
          }else{
            ParseFile(scanDir,onlydata)
          }
}

func ParseFile(fpath string,onlydata *bool){
  if strings.HasSuffix(strings.ToLower(fpath),"png"){
    if(!*onlydata){fmt.Printf("found PNG file %s\n",fpath)}
    imReader,_ := os.Open(fpath)
      thisImg, err := png.Decode(imReader)
      if(err!=nil){
        fmt.Fprintf(os.Stderr,"Error while decoding image %s : %s\n",fpath,err)
          return
      }

bounds := thisImg.Bounds()
          sumR,sumG,sumB := 0., 0., 0.
          totPixels, totNonTransparent := 0, 0
          for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
            for x := bounds.Min.X; x < bounds.Max.X; x++ {
              r,g,b,a := thisImg.At(x, y).RGBA()
                totPixels++
                if a!=65535{
                  //fmt.Printf("found channel alpha %d %d %d %d at %d %d\n",a,r,g,b,x,y)
                  continue
                }
              totNonTransparent++
                sumR+=float64(r)
                sumG+=float64(g)
                sumB+=float64(b)
            }
          }
        if(!*onlydata){
          fmt.Printf("found %d non transparent pixels of %d total, (%d%%)\n",totNonTransparent,totPixels,(totNonTransparent*100)/totPixels)
            fmt.Printf("average RGB values: %f %f %f\n",sumR/float64(totNonTransparent),sumG/float64(totNonTransparent),sumB/float64(totNonTransparent))
        } else
        {
          fmt.Printf("%s\t%f\t%f\t%f\n",fpath,sumR/float64(totNonTransparent),sumG/float64(totNonTransparent),sumB/float64(totNonTransparent))
        }
  }

}
