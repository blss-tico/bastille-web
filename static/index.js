window.onload = () => {
  sessionStorage.setItem('id', crypto.randomUUID());
}

async function sendHttpRequest(url, bodyData) {
  const response = await fetch('/' + url, {
    method: 'POST',
    headers: {
      'Access-Control-Allow-Origin': '*',
      'accept': 'application/json',
      'Content-Type': 'application/json',
      'X-Request-ID': sessionStorage.getItem('id')
    },
    body: bodyData
  });

  if (!response.ok) {    
    const error = await response.text();
    // console.log('index.js - sendHttpRequest - error - response status: ', response.status, error);
    return { code: response.status, error: error };
  }
  
  const responseData = await response.json();
  // console.log('index.js - sendHttpRequest - response status: ', response.status, responseData);
  return { code: response.status, data: responseData };
}
