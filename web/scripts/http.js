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
                        },
                        {
                          label: "HTTP/3",
                          data: data[2],
                          backgroundColor: backgroundColor,
                          borderColor: borderColor,
                          borderWidth: 1
                        }
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


async function fetchDataHTTP() {
        const jsondata = [];
        const response = await fetch('https://raw.githubusercontent.com/derekargueta/irm-data/master/results.csv')

        const data = await response.text()
        const vals = data.split('\n')
        for (let i = 1; i < vals.length-1; i++) {
          const row = vals[i].split(',');
          jsondata.push({
            timestapmp : row[0],
            http2 : row[2],
            http11 : row[3],
            http3 : row[4],
          });
        }

        const timelabel = [];
        const http = [[],[],[],[]];
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
      fetchDataHTTP(); 