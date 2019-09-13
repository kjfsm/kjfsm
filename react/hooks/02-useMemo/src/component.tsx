import React, { useCallback, useMemo, useState } from "react";

const Component: React.FC = () => {
    const [count, setCount] = useState(0); // count と setCount を用意する 初期値は0で setCountを使ってcountの値を更新できるようになる
    const double = useMemo(() => count * 2, [count]); // count が変化するたび再計算が行われる
    const doubleWithUnit = useMemo(() => `${double} pt`, [double]); // double が変化するたび再計算が行われる

    const handlelick = useCallback(() => {
        setCount((prev) => prev + 1);
    }, []);
    return (
        <div>
            <p>count: {count}</p>
            <p>double: {double}</p>
            <p>doubleWithUnit: {doubleWithUnit}</p>
            <button onClick={handlelick}> +1</button>
        </div>
    );
};

export default Component;
