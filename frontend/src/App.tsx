import React, { useState, useEffect } from "react";
import Balloon from "./components/Balloon";
import axios from "axios"


function App() {
  const [balloons, setBalloons] = useState<{text: string, color: string}[]>([]);

  useEffect(() => {

    const interval = setInterval(() => {
      fetch()
    }, 5000);

    const fetch = async() => {
      const res = await axios.get("http://localhost:8000/request")
      setBalloons(prevBalloons => [...prevBalloons, {text: res.data.text, color: res.data.color}])
    }
    return () => clearInterval(interval)
  })

  return (
    <div className="h-screen">
      {balloons.map((balloon, index) => <MemoBalloon key={index} text={balloon.text} color={balloon.color} />)}
    </div>
  );
}

const MemoBalloon = React.memo((props: React.PropsWithChildren<{text: string; color: string;}>) => {
  return <Balloon {...props} />
}, (prevProps, nextProps) => {
  return prevProps.text === nextProps.text && prevProps.color === nextProps.color
});

export default App;
