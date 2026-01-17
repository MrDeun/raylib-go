package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func decodeVertex(chunks []string) Vector3 {
	x, err := strconv.ParseFloat(chunks[1], 64)
	y, err := strconv.ParseFloat(chunks[2], 64)
	z, err := strconv.ParseFloat(chunks[3], 64)

	if err != nil {
		log.Fatal(err.Error())
	}

	return vector3(x, y, z)
}

func decodeFace(chunks []string) [3][3]int {
	coords := [3][3]int{}
	for i := 1; i < len(chunks); i++ {
		vector_chunks := strings.Split(chunks[i], "/")

		vert_coord, err := strconv.ParseInt(vector_chunks[0], 10, 64)
		if err != nil {
			log.Fatal(err.Error())
		}
		coords[i-1][0] = int(vert_coord)
	}
	for _, coor_line := range coords {
		fmt.Printf("%v \n", coor_line)
	}
	return coords
}

func parseOBJFileToGeometry(path string, scale float64) Geometry {
	verts := []Vector3{}
	tris := [][3]*Vector3{}

	f, err := os.Open(path)
	if err != nil {
		log.Fatal(err.Error())
		panic(-1)
	}

	reader := bufio.NewReader(f)

	finished := false
	for !finished {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			finished = true
		} else if err != nil {
			log.Fatal(err.Error())
			panic(-1)
		}

		line_str := string(line)
		chunks := strings.Split(line_str, " ")
		if chunks[0] == "v" {
			vert := decodeVertex(chunks)
			verts = append(verts, vert.MultiplyScalar(scale))
		} else if chunks[0] == "f" {
			_ = decodeFace(chunks)
		}

	}
	return Geometry{position: Vector3Zero(), rotation: Vector3Zero(), verticies: verts, lines: []Line3D{}, triangles: tris}
}
