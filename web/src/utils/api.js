const getSearchResults = async (searchTerm) => {
  const response = await fetch('http://localhost:5000/search', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify({ searchTerm }),
  });
  const data = await response.json();
  return data.results;
};

export { getSearchResults };
