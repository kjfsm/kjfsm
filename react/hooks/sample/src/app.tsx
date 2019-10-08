import React, { useState } from "react";
import Counter from "./counter";

const App: React.FC = () => {
  const [count1, setCount1] = useState(0);
  const [count2, setCount2] = useState(0);
  return (
    <>
      <div>{count1}</div>
      <button onClick={() => setCount1(count1 + 1)}>add1</button>
      <Counter count={count2} setCount={setCount2} />
      <div>合計: {count1 + count2}</div>
    </>
  );
};

export default App;
