import React, { useState, useEffect, useRef } from "react";
import Balloon from "./components/Balloon";
import { w3cwebsocket } from "websocket";


function App() {
  const [balloons, setBalloons] = useState<{text: string, color: string, id:number}[]>([]);

  let counterId = useRef(0)

  useEffect(() => {
    const client = new w3cwebsocket('ws://127.0.0.1:3000/ws', undefined, undefined);

    client.onopen = function() {
      console.log('WebSocket Client Connected');
    };

    client.onmessage = function(message) {
      const data = JSON.parse(message.data.toString());
      console.log(data["Message"])
      setBalloons(prevBalloons => [...prevBalloons, {text: data["Message"], color: data["Color"], id: counterId.current}]);
      counterId.current += 1;
    };

    return () => {
      client.close();
    };
  }, []);

  return (
    <div className="h-screen">
      {balloons.map((balloon, index) => <MemoBalloon key={index} text={balloon.text} color={balloon.color} id={balloon.id}/>)}
    </div>
  );
}

const MemoBalloon = React.memo((props: React.PropsWithChildren<{text: string; color: string; id:number}>) => {
  return <Balloon {...props} />
}, (prevProps, nextProps) => {
  return prevProps.text === nextProps.text && prevProps.color === nextProps.color && prevProps.id === nextProps.id
});

export default App;
