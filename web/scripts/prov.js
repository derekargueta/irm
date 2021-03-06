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
if (chartID == "Chart3") {
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
              plugins: {
                tooltip: {
                  mode: 'index',
                  intersect: false
                }
              },
        scales: {
            y: {
              min: 0,
            max: 100,
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
              plugins: {
                tooltip: {
                  mode: 'index',
                  intersect: false
                }
              },
        scales: {
            y: {
              min: 0,
            max: 100,
                ticks: {
                    callback: function(value, index, values) {
                        return value + '%';
                    }
                }
            }
        }
    }
          });
        } else if (chartID == "Chart5") {
          const myChart = new Chart(ctx, {
            type: 'line',
            data: {
              labels: datelabels,
              datasets: [
                {
                  label: "Stackpath enabled",
                  data: data[0],
                  backgroundColor: backgroundColor,
                  borderColor: borderColor,
                  borderWidth: 1
                },
              ]
            },
            options: {
              plugins: {
                tooltip: {
                  mode: 'index',
                  intersect: false
                }
              },
        scales: {
            y: {
              min: 0,
            max: 100,
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
  async function totalstackpath() {
        const jsondata = [];
        const response = await fetch('https://raw.githubusercontent.com/derekargueta/irm-data/master/results.csv');

        const data = await response.text();
        const vals = data.split('\n');
        for (let i = 1; i < vals.length-1; i++) {
          const row = vals[i].split(',');
          jsondata.push({
            time: row[0],
            stackpath: row[17],
          });
        }

        const stackpath = [[],[]];
        const time = [];
        let i = 0;
        while (jsondata[i] != undefined) {
          time.push(jsondata[i].time);
          stackpath[0].push(jsondata[i].stackpath.replace("%",""));
       
          i++;

        }
        chart(stackpath, time, "% of Stackpath support", "Chart5");
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
            fastly : row[14],
            fastlyipv4 : row[15],
            fastlyipv6 : row[16],
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
      async function fetchdataCloudflare() {
        const jsondata = [];
        const response = await fetch('https://raw.githubusercontent.com/derekargueta/irm-data/master/results.csv');

        const data = await response.text();
        const vals = data.split('\n');
        for (let i = 1; i < vals.length-1; i++) {
          const row = vals[i].split(',');
          jsondata.push({
            time: row[0],
            cloud : row[11],
            cloudipv4 : row[12],
            cloudipv6 : row[13],
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
      totalstackpath(); 
      fetchdataFastly();
      fetchdataCloudflare();