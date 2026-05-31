import React, { useState, useEffect } from "react";

import SearchBar from "./SearchBar";
import { SEARCH_KEY, BASE_URL, TOKEN_KEY } from "../constants";

function Collection(props) {
  const [searchOption, setSearchOption] = useState({
    type: SEARCH_KEY.all,
    keyword: "",
  });

  const handleSearch = (option) => setSearchOption(option);

  return (
    <div className="home">
      <SearchBar handleSearch={handleSearch} />
    </div>
  );
}

export default Collection;
