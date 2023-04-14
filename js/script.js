// Copyright (c) 2020 Mr. Coxall All rights reserved
//
// Created by: Mariam Kasim
// Created on: Apr 2023
// This file contains the JS functions for index.html

'use strict'
/**
 * This function calculates the conversion of temperature from fahrenheit to celsius
 */
function calculate() {
  // input
  const fahrenheit = parseFloat(document.getElementById("fahrenheit").value)

  // process
  const celsius = (fahrenheit - 32) * 5 / 9

  // output
  document.getElementById('celsius').innerHTML = 'celsius is: ' + celsius.toFixed(2) + '°C'
}
