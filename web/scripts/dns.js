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
                  label: "DNS ANY",
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


  async function dnsany() {
        const jsondata = [];
        const response = await fetch('https://raw.githubusercontent.com/derekargueta/irm-data/master/results.csv');

        const data = await response.text();
        const vals = data.split('\n');
        for (let i = 1; i < vals.length-1; i++) {
          const row = vals[i].split(',');
          jsondata.push({
            time: row[0],
            dnsany : row[21]
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
      dnsany(); 