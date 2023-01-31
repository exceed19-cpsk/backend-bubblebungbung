import React, { useEffect } from 'react'
import {randomWithoutMax} from '../functions/utils'
import "../styles/style.scss"

const Balloonv2 = () => {

    const balloonContainer = document.getElementById("balloon-container");

    function getRandomStyles() {
        var r:number = randomWithoutMax(255);
        var g:number = randomWithoutMax(255);
        var b:number = randomWithoutMax(255);
        var mt:number = randomWithoutMax(200);
        var ml:number = randomWithoutMax(50);
        var dur:number = randomWithoutMax(5) + 5;
        return `
        background-color: rgba(${r},${g},${b},0.7);
        color: rgba(${r},${g},${b},0.7); 
        box-shadow: inset -7px -3px 10px rgba(${r - 10},${g - 10},${b - 10},0.7);
        margin: ${mt}px 0 0 ${ml}px;
        animation: float ${dur}s ease-in infinite
        `;
      }

    function createBalloons(num:number) {
        for (var i = num; i > 0; i--) {
          var balloon = document.createElement("div");
          balloon.className = "balloon";
          balloon.style.cssText = getRandomStyles();
          balloonContainer?.append(balloon);
        }
      }

      useEffect(() => {
        createBalloons(30)
      })

  return (
    <div id="balloon-container">
    </div>
  )
}

export default Balloonv2