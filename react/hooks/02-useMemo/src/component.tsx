import React, { useCallback, useMemo, useState } from "react";

const Component: React.FC = () => {
    const [count, setCount] = useState(0);
    const double = useMemo(() => count * 2, [count]);
    const doubleWithUnit = useMemo(() => `${double} pt`, [double]);

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
