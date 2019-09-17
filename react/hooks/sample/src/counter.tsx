import React from "react";

interface CounterProps {
    count: number;
    setCount: React.Dispatch<number>;
}

const Counter: React.FC<CounterProps> = (props) => {
    return (
        <>
            <div>{props.count}</div>
            <button onClick={() => props.setCount(props.count + 1)}>add2</button>
        </>
    );
};
export default Counter;
