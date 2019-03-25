package main

import(

)
type position struct{
	x,t float64
}
type entity struct{
	pos position
	rotation float64
	active bool
}
