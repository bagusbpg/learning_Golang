/*
There is a shop with old-style cash registers. Rather than scanning itmes and pulling
the price from a database, the price of each item is typed in manually. This method
sometimes leads to errors. Given a list of items and their correct prices, compare
the prices to those enterd when each item was sold. Determine the number of errors in
selling prices.

Example
products = ['eggs', 'milk', 'cheese']
productPrices = [2.89, 3.29, 5.79]
productSold = ['eggs', 'eggs', 'cheese', 'milk']
soldPrice = [2.89, 2.99, 5.97, 3.29]

          Price
Product   Actual   Expected    Error
eggs      2.89     2.89
eggs      2.99     2.89        1
cheese    5.97     5.79        1
milk      3.29     3.29

The second sale of eggs has a wrong price, as does the sale of cheese. There are 2 errors
in pricing.

Function Description
Complete the function priceCheck in the editor below.
priceCheck has the following parameter(s):
string products[n]: each products[i] is the name of an item for sale
float productPrices[n]: each productPrices[i] is the price of products[i]
string productSold[m]: each productSold[j] is the name of a product sold
float soldPrice[m]: each soldPrice[j] contains the sale price recorded for productSold[j]

Returns
int: the number of sale prices that were entered incorrectly

Constraints
1 <= n <= 10^5
1 <= m <= n
1.00 <= productPrices[i], soldPrice[j] <= 100000.00, where 0 <= i < n, and 0 <= j < m
*/

package main

import "fmt"

func priceCheck(products []string, productPrices []float32, productSold []string, soldPrice []float32) int32 {
	// Write your code here
	var mapProduct map[string]float32 = map[string]float32{}
	var count int32 = 0

	for i := 0; i < len(products); i++ {
		if _, exist := mapProduct[products[i]]; !exist {
			mapProduct[products[i]] = productPrices[i]
		}
	}

	for i := 0; i < len(productSold); i++ {
		if mapProduct[productSold[i]] != soldPrice[i] {
			count++
		}
	}

	return count
}

func main() {
	var n, m int
	var name string
	var price float32
	var products, productSold []string = []string{}, []string{}
	var productPrices, soldPrice []float32 = []float32{}, []float32{}

	fmt.Scanf("%v\n", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &name)
		products = append(products, name)
	}
	for i := 0; i < n; i++ {
		fmt.Scanf("%v\n", &price)
		productPrices = append(productPrices, price)
	}

	fmt.Scanf("%v\n", &m)
	for i := 0; i < m; i++ {
		fmt.Scanf("%v\n", &name)
		productSold = append(productSold, name)
	}
	for i := 0; i < m; i++ {
		fmt.Scanf("%v\n", &price)
		soldPrice = append(soldPrice, price)
	}

	fmt.Println(priceCheck(products, productPrices, productSold, soldPrice))
}
