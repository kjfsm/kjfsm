import React from 'react';

interface Props {
    value: string;
    onClick: (i: number) => void;
}

class Square extends React.Component<Props> {
 
    onClick = () => {
        this.props.onClick()
    }

    render = () => {
        return (
            <button className="square" onClick={this.onClick} >
                {this.props.value}
            </button>
        );
    }
}
export default Square;