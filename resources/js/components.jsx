import React, {useState, useEffect} from "react";

export const Header = () => {
    return (<h1 className="text-3xl font-bold">test in header</h1>);
};

export const Body = () => {
    return (
        <>
        <h1>Test</h1>
        <Test />
        </>
    );
};

export const Test = () => {
    const [counter, setCounter] = useState(0);        

    useEffect(() => {
        console.log(counter);
    }, [counter]);

    return (
        <>
            Counter: {counter}

            <button onClick={() => setCounter((prev) => prev + 1)}>Increment</button>
            <h1 className="text-xl">React component Header</h1>
        </>
    );
}
