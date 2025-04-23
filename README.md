#### Setup

- go mod tidy

#### Run Webhook

- go run main.go webhook --port[optional] 8080

#### Set Max Customer per Agent

1. Start server : go run main.go server --port[optional] 8081
2. Setter : curl -X PUT https://40f2-103-102-12-15.ngrok-free.app/config -H "Content-Type: application/json" -d '{"max_customer_per_agent": 2}'
3. Getter : curl -X GET https://40f2-103-102-12-15.ngrok-free.app/config

#### Run Dequeue Listener

- go run main.go webhook

#### Run Resolver Listener

- go run main.go resolve

#### Start Chat

- Please paste the script inside `<body>` tag `</body>`

```
<script>
    document.addEventListener('DOMContentLoaded', function() {
        var s,t; s = document.createElement('script'); s.type = 'text/javascript';
        s.src = 'https://omnichannel.qiscus.com/js/qismo-v4.js'; s.async = true;
        s.onload = s.onreadystatechange = function() { new Qismo('rvcbl-fcsngqk40iyo7ks', {
                        options: {
                            channel_id: 130821, qismoIframeUrl: 'https://omnichannel.qiscus.com', baseUrl: 'https://omnichannel.qiscus.com',
                            extra_fields: [],
                        }
                    }); }
        t = document.getElementsByTagName('script')[0]; t.parentNode.insertBefore(s, t);
    });
</script>
```

#### Technical Documentation

- Flowchart
  [![3GtEAS1.th.png](https://iili.io/3GtEAS1.th.png)](https://freeimage.host/i/3GtEAS1)

- Sequence Diagram
  [![3GtMwG4.th.png](https://iili.io/3GtMwG4.th.png)](https://freeimage.host/i/3GtMwG4)

- Database Design
  [![3GthszF.th.png](https://iili.io/3GthszF.th.png)](https://freeimage.host/i/3GthszF)
