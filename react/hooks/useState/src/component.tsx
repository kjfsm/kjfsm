import React, { useCallback, useState } from "react";

const Component: React.FC = () => {
    const [count, setCount] = useState(0);
    const handlelick = useCallback(() => {
        setCount(count + 1);
    }, [count]);
    return(
        <div>
            <p>{count}</p>
            <button onClick={handlelick}> +1</button>
        </div>
    );
};

export default Component;
