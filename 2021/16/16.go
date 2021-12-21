package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Packet struct {
	Data         string
	Type         PacketType
	LengthTypeId int
	Value        int64
	Version      int
	Children     []*Packet
}

type PacketType int

const (
	TYPE_SUM PacketType = iota
	TYPE_PROD
	TYPE_MIN
	TYPE_MAX
	TYPE_LITERAL
	TYPE_GREATERTHAN
	TYPE_LESSERTHAN
	TYPE_EQUALTO
)

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
	versionValue, _ := strconv.ParseInt(p.Data[:3], 2, 32)
	typeValue, _ := strconv.ParseInt(p.Data[3:6], 2, 32)

	p.Version = int(versionValue)
	p.Type = PacketType(typeValue)
}

func (p *Packet) parsePacket() int {
	// Parse version and type
	p.parsePacketHeader()

	// Then act on literal or operator packet
	if p.Type == TYPE_LITERAL {
		return p.parseLiteralPacket()
	} else {
		n := p.parseOperatorPacket()
		p.evaluate()
		return n
	}
}

func (p *Packet) parseLiteralPacket() int {
	bitFields := []string{}

	var end int
	for i := 6; i < len(p.Data); i += 5 {
		chunk := p.Data[i : i+5]
		bitFields = append(bitFields, chunk[1:])

		if strings.HasPrefix(chunk, "0") {
			end = i + 5
			break
		}
	}

	fullBitField := strings.Join(bitFields, "")
	v, _ := strconv.ParseInt(fullBitField, 2, 64)

	p.Value = v
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
	return used
}

func (p *Packet) evaluate() {
	switch p.Type {
	case TYPE_SUM: // Sum
		var sum int64
		for _, child := range p.Children {
			sum += child.Value
		}
		p.Value = sum
	case TYPE_PROD: // Prod
		prod := int64(1)
		for _, child := range p.Children {
			prod *= child.Value
		}
		p.Value = prod
	case TYPE_MIN: // Min
		min := int64(math.MaxInt64)
		for _, child := range p.Children {
			if min > child.Value {
				min = child.Value
			}
		}
		p.Value = min
	case TYPE_MAX: // Max
		var max int64
		for _, child := range p.Children {
			if max < child.Value {
				max = child.Value
			}
		}
		p.Value = max
	case TYPE_GREATERTHAN: // GT
		if p.Children[0].Value > p.Children[1].Value {
			p.Value = 1
		} else {
			p.Value = 0
		}
	case TYPE_LESSERTHAN: // LT
		if p.Children[0].Value < p.Children[1].Value {
			p.Value = 1
		} else {
			p.Value = 0
		}
	case TYPE_EQUALTO: // EQ
		if p.Children[0].Value == p.Children[1].Value {
			p.Value = 1
		} else {
			p.Value = 0
		}
	}
}

func sumVersions(p *Packet) int64 {

	var sum int64
	for _, child := range p.Children {
		sum += sumVersions(child)
	}

	return int64(p.Version) + sum
}

// func debugPartTwo(p *Packet) {
// 	log.Printf("Root - V: %d T: %d Val: %d", p.Version, p.Type, p.Value)
// 	for _, packet := range p.Children {
// 		log.Printf("V: %d T: %d Val: %d", packet.Version, packet.Type, packet.Value)
// 	}
// }

func run(input string, returnValue bool) int64 {
	message := hexToBinaryString(input)

	p := Packet{
		Data: message,
	}
	p.parsePacket()
	p.evaluate()

	if !returnValue {
		return sumVersions(&p)
	} else {
		// debugPartTwo(&p)
		return int64(p.Value)
	}
}
