package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Packet struct {
	Data         string
	Type         int
	LengthTypeId int
	Value        int
	Version      int
	Children     []*Packet
}

func hexToBinaryString(input string) string {

	var binary []string
	for _, chr := range input {
		v, _ := strconv.ParseInt(string(chr), 16, 32)
		s := fmt.Sprintf("%04s", strconv.FormatInt(v, 2))

		binary = append(binary, s)
	}
	return strings.Join(binary, "")
}

func (p *Packet) parsePacketHeader() {
	version := p.Data[:3]
	typeId := p.Data[3:6]

	versionValue, _ := strconv.ParseInt(version, 2, 32)
	typeValue, _ := strconv.ParseInt(typeId, 2, 32)

	p.Version = int(versionValue)
	p.Type = int(typeValue)
}

func (p *Packet) parsePacket() int {
	// If only 000s just return and provide value
	//  to eat rest of packet bitfield
	if strings.Count(p.Data, "0") == len(p.Data) {
		return len(p.Data)
	}

	// Parse version and type
	p.parsePacketHeader()

	// Then act on literal or operator packet
	if p.Type == 4 {
		return p.parseLiteralPacket()
	} else {
		return p.parseOperatorPacket()
	}
}

func (p *Packet) parseLiteralPacket() int {
	bitFields := []string{}

	var end int
	for i := 6; i < len(p.Data); i += 5 {
		chunk := p.Data[i : i+5]
		bitFields = append(bitFields, chunk)

		if strings.HasPrefix(chunk, "0") {
			end = i + 5
			break
		}
	}

	v, _ := strconv.ParseInt(strings.Join(bitFields, ""), 2, 32)

	p.Value = int(v)
	return end
}

func (p *Packet) parseOperatorPacket() int {
	v, _ := strconv.ParseInt(string(p.Data[6]), 2, 32)
	p.LengthTypeId = int(v)

	var used int
	if p.LengthTypeId == 0 {
		subPacketTotalBitLength, _ := strconv.ParseInt(p.Data[7:22], 2, 32)

		used = 22
		tbl := int(subPacketTotalBitLength) + used
		for used < tbl {
			newP := Packet{
				Data: p.Data[used:],
			}
			p.Children = append(p.Children, &newP)
			used += newP.parsePacket()
		}
	} else if p.LengthTypeId == 1 {
		nSubPackets, _ := strconv.ParseInt(p.Data[7:18], 2, 32)

		nPackets := 1
		used = 18
		for nPackets <= int(nSubPackets) {

			newP := Packet{
				Data: p.Data[used:],
			}
			p.Children = append(p.Children, &newP)
			used += newP.parsePacket()
			nPackets++
		}
	}

	switch p.Type {
	case 0:
		var sum int
		for _, child := range p.Children {
			sum += child.Value
		}
		p.Value = sum
	case 1:
		prod := 1
		for _, child := range p.Children {
			prod *= child.Value
		}
		p.Value = prod
	case 2:
		min := math.MaxInt
		for _, child := range p.Children {
			if min > child.Value {
				min = child.Value
			}
		}
		p.Value = min
	case 3:
		var max int
		for _, child := range p.Children {
			if max < child.Value {
				max = child.Value
			}
		}
		p.Value = max
	case 5:
		if p.Children[0].Value > p.Children[1].Value {
			p.Value = 1
		} else {
			p.Value = 0
		}
	case 6:
		if p.Children[0].Value < p.Children[1].Value {
			p.Value = 1
		} else {
			p.Value = 0
		}
	case 7:
		if p.Children[0].Value == p.Children[1].Value {
			p.Value = 1
		} else {
			p.Value = 0
		}
	}

	return used
}

func sumVersions(p *Packet) int {

	var sum int
	for _, child := range p.Children {
		sum += sumVersions(child)
	}

	return p.Version + sum
}

func run(input string, returnValue bool) int {
	message := hexToBinaryString(input)

	p := Packet{
		Data: message,
	}
	p.parsePacket()

	if !returnValue {
		return sumVersions(&p)
	} else {
		return p.Value
	}
}
