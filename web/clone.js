function chart(data, datelabels, minilabel, chartID) {
  const ctx = document.getElementById(chartID).getContext('2d');

  // All the charts have the same color schemes.
  const backgroundColor = [
    'rgba(255, 99, 132, 0.2)',
    'rgba(54, 162, 235, 0.2)',
    'rgba(255, 206, 86, 0.2)',
    'rgba(75, 192, 192, 0.2)',
    'rgba(153, 102, 255, 0.2)',
    'rgba(255, 159, 64, 0.2)'
  ];
  const borderColor = [
    'rgba(255, 99, 132, 1)',
    'rgba(54, 162, 235, 1)',
    'rgba(255, 206, 86, 1)',
    'rgba(75, 192, 192, 1)',
    'rgba(153, 102, 255, 1)',
    'rgba(255, 159, 64, 1)'
  ];
  const options = {
    scales: {
      y: {
        beginAtZero: true
      }
    }
  };

  if (chartID == "Chart1") {
    const myChart = new Chart(ctx, {
      type: 'line',
      data: {
        labels: datelabels,
        datasets: [
          {
            label: "HTTP/1.1",
            data: data[0],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          },
          {
            label: "HTTP/2",
            data: data[1],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          }, {
            label: "HTTP/3",
            data: data[2],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          }
        ]
      },
      options: {
  scales: {
      y: {
          ticks: {
              callback: function(value, index, values) {
                  return value + '%';
              }
          }
      }
  }
}
    });
  } else if (chartID == "Chart2") {
    const myChart = new Chart(ctx, {
      type: 'line',
      data: {
        labels: datelabels,
        datasets: [
          {
            label: "TLSv1.0",
            data: data[0],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          },
          {
            label: "TLSv1.1",
            data: data[1],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          },
          {
            label: "TLSv1.2",
            data: data[2],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          },
          {
            label: "TLSv1.3",
            data: data[3],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          }
        ]
      },
      options: {
  scales: {
      y: {
          ticks: {
              callback: function(value, index, values) {
                  return value + '%';
              }
          }
      }
  }
}
    });
  }else if (chartID == "Chart3") {
    const myChart = new Chart(ctx, {
      type: 'line',
      data: {
        labels: datelabels,
        datasets: [
          {
            label: "Cloudflare Total",
            data: data[0],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          },
          {
            label: "Cloudflare ipv4",
            data: data[1],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          },
          {
            label: "Cloudflare ipv6",
            data: data[2],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          },
          
        ]
      },
      options: {
  scales: {
      y: {
          ticks: {
              callback: function(value, index, values) {
                  return value + '%';
              }
          }
      }
  }
}
    });
  }else if (chartID == "Chart4") {
    const myChart = new Chart(ctx, {
      type: 'line',
      data: {
        labels: datelabels,
        datasets: [
          {
            label: "Fastly Total",
            data: data[0],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          },
          {
            label: "Fastly ipv4",
            data: data[1],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          },
          {
            label: "Fastly ipv6",
            data: data[2],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          },
          
        ]
      },
      options: {
  scales: {
      y: {
          ticks: {
              callback: function(value, index, values) {
                  return value + '%';
              }
          }
      }
  }
}
    });
  }else if (chartID == "Chart5") {
    const myChart = new Chart(ctx, {
      type: 'line',
      data: {
        labels: datelabels,
        datasets: [
          {
            label: "Dualstack",
            data: data[0],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          },
          {
            label: "Ipv4",
            data: data[1],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          },
          {
            label: "Ipv6",
            data: data[2],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          },
          
        ]
      },
      options: {
  scales: {
      y: {
          ticks: {
              callback: function(value, index, values) {
                  return value + '%';
              }
          }
      }
  }
}
    });
  }else if (chartID == "Chart6") {
    const myChart = new Chart(ctx, {
      type: 'line',
      data: {
        labels: datelabels,
        datasets: [
          {
            label: "DNS ANY",
            data: data[0],
            backgroundColor: backgroundColor,
            borderColor: borderColor,
            borderWidth: 1
          }, 
        ]
      },
      options: {
  scales: {
      y: {
          ticks: {
              callback: function(value, index, values) {
                  return value + '%';
              }
          }
      }
  }
}
    });
  }
}

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
      http3: row[4],
      
    });
  }

  const timelabel = [];
  const http = [[],[],[]];
  let i = 0;
  while (jsondata[i] != undefined) {
    timelabel.push(jsondata[i].timestapmp);
    http[0].push(jsondata[i].http11.replace("%",""));
    http[1].push(jsondata[i].http2.replace("%",""));
    http[2].push(jsondata[i].http3.replace("%",""));
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
async function fetchdataCloudflare() {
  const jsondata = [];
  const response = await fetch('https://raw.githubusercontent.com/derekargueta/irm-data/master/results.csv');

  const data = await response.text();
  const vals = data.split('\n');
  for (let i = 1; i < vals.length-1; i++) {
    const row = vals[i].split(',');
    jsondata.push({
      time: row[0],
      cloud : row[10],
      cloudipv4 : row[11],
      cloudipv6 : row[12],
    });
  }

  const cloud = [[],[],[]];
  const time = [];
  let i = 0;
  while (jsondata[i] != undefined) {
    time.push(jsondata[i].time);
    cloud[0].push(jsondata[i].cloud.replace("%",""));
    cloud[1].push(jsondata[i].cloudipv4.replace("%",""));
    cloud[2].push(jsondata[i].cloudipv6.replace("%",""));
    i++;
  }

  chart(cloud, time, "% of Cloudflare support", "Chart3");
}

async function fetchdataFastly() {
  const jsondata = [];
  const response = await fetch('https://raw.githubusercontent.com/derekargueta/irm-data/master/results.csv');

  const data = await response.text();
  const vals = data.split('\n');
  for (let i = 1; i < vals.length-1; i++) {
    const row = vals[i].split(',');
    jsondata.push({
      time: row[0],
      fastly : row[13],
      fastlyipv4 : row[14],
      fastlyipv6 : row[15],
    });
  }

  const fastly = [[],[],[]];
  const time = [];
  let i = 0;
  while (jsondata[i] != undefined) {
    time.push(jsondata[i].time);
    fastly[0].push(jsondata[i].fastly.replace("%",""));
    fastly[1].push(jsondata[i].fastlyipv4.replace("%",""));
    fastly[2].push(jsondata[i].fastlyipv6.replace("%",""));
    i++;
  }

  chart(fastly, time, "% of Cloudflare support", "Chart4");
}


async function totalipv() {
  const jsondata = [];
  const response = await fetch('https://raw.githubusercontent.com/derekargueta/irm-data/master/results.csv');

  const data = await response.text();
  const vals = data.split('\n');
  for (let i = 1; i < vals.length-1; i++) {
    const row = vals[i].split(',');
    jsondata.push({
      time: row[0],
      ipvtotal : row[16],
      ipv4 : row[17],
      ipv6 : row[18],
    });
  }

  const ipv = [[],[],[]];
  const time = [];
  let i = 0;
  while (jsondata[i] != undefined) {
    time.push(jsondata[i].time);
    ipv[0].push(jsondata[i].ipvtotal.replace("%",""));
    ipv[1].push(jsondata[i].ipv4.replace("%",""));
    ipv[2].push(jsondata[i].ipv6.replace("%",""));
    i++;

  }
  chart(ipv, time, "% of IPV support", "Chart5");
}
async function dnsany() {
  const jsondata = [];
  const response = await fetch('https://raw.githubusercontent.com/derekargueta/irm-data/master/results.csv');

  const data = await response.text();
  const vals = data.split('\n');
  for (let i = 1; i < vals.length-1; i++) {
    const row = vals[i].split(',');
    jsondata.push({
      time: row[0],
      dnsany : row[19]
    });
  }

  const dnsany = [[],[]];
  const time = [];
  let i = 0;
  while (jsondata[i] != undefined) {
    time.push(jsondata[i].time);
    dnsany[0].push(jsondata[i].dnsany.replace("%",""));
    i++;

  }
  chart(dnsany, time, "% of DNS ANY support", "Chart6");
}
