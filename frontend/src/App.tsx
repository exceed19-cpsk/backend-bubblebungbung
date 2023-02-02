import { useState, useEffect} from "react";
import Balloon from "./components/Balloon";
import axios from "axios"


function App() {
  const [balloons, setBalloons] = useState<{text: string, color: string}[]>([]);

  useEffect(() => {

    const interval = setInterval(() => {
      fetch()
    }, 3000);

    const fetch = async() => {
      const res = await axios.get("http://localhost:8000/request")
      setBalloons(prevBalloons => [...prevBalloons, {text: res.data.text, color: "red"}])
    }
    return () => clearInterval(interval)
  })

  return (
    <div className="h-screen">
      {balloons.map((balloon, index) => <Balloon key={index} text={balloon.text} color={balloon.color} />)}
    </div>
  );
}

export default App;
