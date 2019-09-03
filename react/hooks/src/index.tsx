import React from "react";
import { render } from "react-dom";
import Component from "./component";

const App: React.FC = () => (
 <div>Hello World App</div>
);

render(<Component />, document.getElementById("root"));
