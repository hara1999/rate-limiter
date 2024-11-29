<h2 align="center">Rate Warden - gRPC based Rate Limiting Service</h1>
    <p><strong>Rate Warden</strong> is a gRPC rate-limiting service designed to provide an easy interface for defining and using any existing rate-limiting algorithm with Redis. Rate Warden supports seamless Kubernetes deployment via Helm Charts.</p>

<h3>Features</h3>
<ul>
    <li><strong>VerifyLimit:</strong> verifies if the application adheres to the registered rate limits.</li>
    <li><strong>RegisterClient:</strong> Registers a new client with custom limits into the Rate Warden database.</li>
    <li><strong>GetClient:</strong> Retrieves existing client limits from the Rate Warden database.</li>
    <li><strong>UpdateClient:</strong> Updates an existing client's rate limits.</li>
    <li><strong>DeleteClient:</strong> Deletes the client.</li>
</ul>
<p>Rate Warden minimizes database I/O calls by using Redis for rate-limiting checks. When a new client is registered or an existing one is updated, the cache is updated with the client limits. Therefore, <b>VerifyLimit</b> does not need to fetch client limits from the database. This also means that rate limits for a client can be updated on the fly since they are maintained as a database entry in Rate Warden.</p>

<h3>Fixed Window Algorithm</h3>
<p>A rate-limiting strategy where, once the request limit is reached, a client is blocked from making further requests until the expiration time elapses. For example, if a client has a limit of 50 requests per minute and uses all 50 requests within the first 5 seconds, it will be blocked from making additional requests for the remaining 55 seconds.</p>
<p>However, this approach is susceptible to bursty traffic patterns. Allowing a client to rapidly exhaust its entire limit can potentially overload the service. In scenarios where traffic is concentrated within a short timeframe, this could lead to unexpected spikes in load, disrupting the anticipated distribution of requests.</p>
<p>While the fixed window algorithm is straightforward and easy to implement, its limitation in handling bursty traffic necessitates careful consideration for managing and balancing the service load. Alternative rate-limiting algorithms might better accommodate fluctuating or uneven request patterns.</p>

<h3>Rolling Window Algorithm</h3>
<p>The rolling window strategy maintains a historical record of traffic within a defined time window, unlike the fixed window strategy. It continuously tracks the traffic generated over a specified period (e.g. the last 5 minutes) to determine if a client should be allowed further requests, offering better protection against sudden traffic bursts.</p>
<p>This approach requires higher CPU and memory resources due to the need to store more request information with timestamps. However, it improves resilience against traffic spikes by evenly distributing requests over time.</p>
<p>Using a sorted set with timestamps allows for efficient removal of expired requests. Employing UUIDs minimizes conflicts within the set, ensuring smooth operation even under heavy traffic loads.</p>

<h3>Setup Guide</h3>
<ol>
    <li>Copy the <code>.env.example</code> to <code>.env</code>.</li>
    <li>Run <code>make dev</code> (assuming <a href="https://sp21.datastructur.es/materials/guides/make-install.html">make</a> and <a href="https://docs.docker.com/compose/install/">docker-compose</a> are installed).</li>
    <li>Install <a href="https://kreya.app/">Kreya</a> for making the gRPC calls.</li>
    <li>Import the proto file into Kreya.</li>
</ol>