window.onload = () => {
  sessionStorage.removeItem('id');
  sessionStorage.removeItem('jails');

  const uuidVal = uuidv7Generate();
  const uuidStr = Array.from(uuidVal).map((b) => b.toString(16).padStart(2, "0")).join("");
  const uuidFmt = uuidv7StringFormat(uuidStr);
  sessionStorage.setItem('id', uuidFmt);

  setTimeout(async () => {
    const formJson = JSON.stringify({"action":"all","options":"-j"});
    const jails = await sendHttpRequest("list", formJson);
    if (jails.data.msg) { sessionStorage.setItem('jails', jails.data.msg); }
  }, 500);
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
    //console.log('index.js - sendHttpRequest - error - response status: ', response.status, error);
    return { code: response.status, error: error };
  }
  
  const responseData = await response.json();
  //console.log('index.js - sendHttpRequest - response status: ', response.status, responseData);
  return { code: response.status, data: responseData };
}

function uuidv7Generate() {
  const value = new Uint8Array(16);
  crypto.getRandomValues(value);

  const timestamp = BigInt(Date.now());
  value[0] = Number((timestamp >> 40n) & 0xffn);
  value[1] = Number((timestamp >> 32n) & 0xffn);
  value[2] = Number((timestamp >> 24n) & 0xffn);
  value[3] = Number((timestamp >> 16n) & 0xffn);
  value[4] = Number((timestamp >> 8n) & 0xffn);
  value[5] = Number(timestamp & 0xffn);
  value[6] = (value[6] & 0x0f) | 0x70;
  value[8] = (value[8] & 0x3f) | 0x80;

  return value;
}

function uuidv7StringFormat(uuid) {
  const p1 = uuid.slice(0, 8);
  const p2 = uuid.slice(8, 12);
  const p3 = uuid.slice(12, 16);
  const p4 = uuid.slice(16, 20);
  const p5 = uuid.slice(20);
  const uuidFmt = p1 + "-" + p2 + "-" + p3 + "-" + p4 + "-" + p5;
  
  return uuidFmt;
}
