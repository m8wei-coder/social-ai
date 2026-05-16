import React, { useState } from "react";
import Main from "./Main";

import { TOKEN_KEY } from "../constants";

const App = (props) => {
    // TODO: check if token is expired
    // option1: ttl or exp time
    // option2: call backend to check if token is valid
    const [isLoggedIn, setIsLoggedIn] = useState(
      localStorage.getItem(TOKEN_KEY) ? true : false
    );

    const loggedIn = (token) => {
      if (token) {
        localStorage.setItem(TOKEN_KEY, token);
        setIsLoggedIn(true);
      }

    }
    return <div className="App">
      <Main isLoggedIn={isLoggedIn} handleLoggedIn={loggedIn} />
    </div>;
}

export default App
