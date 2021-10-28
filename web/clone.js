async function fetchDataHTTP() {
    const jsondata = [];
    const response = await fetch('https://raw.githubusercontent.com/derekargueta/irm-data/master/results.csv')

    const data = await response.text()
    const vals = data.split('\n')
    for (let i = 1; i < vals.length-1; i++) {
      const row = vals[i].split(',');
      jsondata.push({
        timestapmp : row[0],
        domaintest : row[1],
        http2 : row[2],
        http11 : row[3],
      });
    }

    const timelabel = [];
    const http = [[],[]];
    let i = 0;
    while (jsondata[i] != undefined) {
      timelabel.push(jsondata[i].timestapmp);
      http[0].push(jsondata[i].http11.replace("%",""));
      http[1].push(jsondata[i].http2.replace("%",""));
      i++;
    }

    chart(http, timelabel, "% of HTTP supported Websites", "Chart1");
  }
  
  async function fetchDataTLS() {
    const jsondata = [];
    const response = await fetch('https://raw.githubusercontent.com/derekargueta/irm-data/master/results.csv');

    const data = await response.text();
    const vals = data.split('\n');
    for (let i = 1; i < vals.length-1; i++) {
      const row = vals[i].split(',');
      jsondata.push({
        time: row[0],
        tls10 : row[6],
        tls11 : row[7],
        tls12 : row[8],
        tls13 : row[9],
      });
    }

    const tls = [[],[],[],[]];
    const time = [];
    let i = 0;
    while (jsondata[i] != undefined) {
      time.push(jsondata[i].time);
      tls[0].push(jsondata[i].tls10.replace("%",""));
      tls[1].push(jsondata[i].tls11.replace("%",""));
      tls[2].push(jsondata[i].tls12.replace("%",""));
      tls[3].push(jsondata[i].tls13.replace("%",""));
      i++;
    }

    chart(tls, time, "% of TLS Websites", "Chart2");
  }