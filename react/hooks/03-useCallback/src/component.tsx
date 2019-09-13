import React, { useCallback, useState } from "react";

interface Props {
    clickedX: number;
    clickedY: number;
    handleClick: (event: React.MouseEvent<HTMLElement, MouseEvent>) => void;
}

const Component: React.FC<Props> = (props) => (
    <div>
        <div
            style={{ width: 100, height: 100, background: "#ccf" }}
            onClick={props.handleClick}
        />
        <p
            style={{ width: 100, height: 100, background: "#ccf" }}
            onClick={props.handleClick}
        />
        <p>X: {props.clickedX}</p>
        <p>Y: {props.clickedY}</p>
    </div>
);

const Container: React.FC = () => {
    const [state, update] = useState({
        clickedX: 0,
        clickedY: 0,
    });

    // イベントを useCallbackでくるむ感じ？
    const handleClick = useCallback((event: React.MouseEvent<HTMLElement, MouseEvent>) => {
        event.persist();
        const { top, left } = event.currentTarget.getBoundingClientRect(); // オブジェクトの左上の角からの座標
        update((prev) => ({
            ...prev,
            clickedX: event.clientX - left, // event.clientX は ブラウザウィンドウの左上の角からの座標
            clickedY: event.clientY - top,
        }));
    },
        [],
    );
    return (
        <Component
            clickedX={state.clickedX}
            clickedY={state.clickedY}
            handleClick={handleClick}
        />
    );
};

export default Container;
