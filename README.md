Go RGB average value
====================

This is an utility written in Go to retrieve the average value of RGB channels of one or more PNG files.

Once compiled it has no dependencies and is made specifically to be called by a shell script.

It takes into account only the non trasparent (100% opaque) pixels

Usage
-----

For human readers:

`./rgb_average folder_path`
`./rgb_average specific_image_path`

For scripts:

The *-t* flag displayes only the filename and the RGB values delimited by tabs

`./rgb_average -t folder_path`
`./rgb_average -t specific_image_path`

This shows the filename and the sum of the three averages of the images in the Download folder:

`./rgb_average -t ../Downloads | awk '//{print $1 " " $2+$3+$4}'`

