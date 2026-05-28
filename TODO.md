# Cloudflare IP Scanner for v2Ray – TODO List

## 1. Preparation
- [ ] Get Cloudflare IPv4 ranges (`curl -s https://www.cloudflare.com/ips-v4`)
- [ ] Get Cloudflare IPv6 ranges (optional)
- [ ] Decide on scan scope (e.g., only /24 subnets, all IPs, or random sample)
- [ ] Choose programming language (Python, Go, Bash + tools)
- [ ] Set rate limits (delay between probes to avoid ban)
- [ ] Prepare output file format (CSV/JSON with IP, ping, download, upload, etc.)

## 2. IP Generation & Pre-filtering
- [ ] Expand CIDR ranges into individual IPs (or use /24 chunks)
- [ ] Remove duplicate IPs
- [ ] Exclude reserved/private IPs (already none in CF ranges)
- [ ] Randomize IP order to avoid scanning sequentially (reduce detection)
- [ ] Optional: keep only IPs that are not already known as bad

## 3. Ping Test (ICMP)
- [ ] Implement ICMP ping with timeout (e.g., 1 second)
- [ ] Send 1–3 pings per IP
- [ ] Calculate packet loss and average RTT
- [ ] Filter: keep only IPs with loss < 20% and RTT < 300 ms
- [ ] Save results with RTT and loss percentage
- [ ] Handle ICMP blocking (fallback to TCP ping later)

## 4. TCP Ping (if ICMP is blocked)
- [ ] Implement TCP SYN to port 443 (or 80)
- [ ] Measure connection time (RTT)
- [ ] Timeout: 1–2 seconds
- [ ] Filter: keep IPs with successful TCP handshake

## 5. Port & Protocol Check
- [ ] Test common Cloudflare ports: 443, 8443, 2053, 2083, 2087, 2096
- [ ] Check if port is open (TCP connect)
- [ ] Verify TLS handshake (if HTTPS port)
- [ ] Record fastest open port per IP

## 6. HTTP/HTTPS Validation (Cloudflare fingerprint)
- [ ] Send GET request with a valid Host header (your proxied domain)
- [ ] Check for Cloudflare-specific headers:
  - `Server: cloudflare`
  - `cf-ray` present
  - `CF-Cache-Status`
- [ ] Accept HTTP status codes: 200, 403, 503 (but not 404 or 500)
- [ ] Measure time to first byte (TTFB)
- [ ] Filter: keep only IPs that respond as Cloudflare

## 7. Latency Refinement
- [ ] Run 5–10 parallel HTTPS requests to calculate average latency
- [ ] Use same Host and path (e.g., `/speedtest/10mb.bin`)
- [ ] Sort IPs by latency (lowest first)

## 8. Download Speed Test
- [ ] Download a small file (1 MB – 10 MB) from behind Cloudflare
- [ ] Measure total download time and speed (MB/s or Mbps)
- [ ] Use `curl` or custom HTTP client with timeout (5–10 sec)
- [ ] Repeat 2–3 times and take average
- [ ] Save download speed per IP

## 9. Upload Speed Test
- [ ] Upload small random data (e.g., 500 KB – 5 MB) via POST
- [ ] Measure upload speed
- [ ] Ensure same path and Host as download test
- [ ] Save upload speed per IP

## 10. Aggregation & Scoring
- [ ] Combine results: IP, ping (ms), download (Mbps), upload (Mbps)
- [ ] Calculate a score (e.g., `(download + upload) / (ping + 1)`)
- [ ] Rank IPs by score
- [ ] Flag unstable IPs (high variance in speed/latency)

## 11. v2Ray Integration Test (Optional but recommended)
- [ ] Use top 5–10 IPs in your actual v2Ray config
- [ ] Test with your real WebSocket + TLS + path settings
- [ ] Measure real proxy speed (throughput, stability)
- [ ] Validate for at least 5 minutes per IP
- [ ] Output final “working” list with notes

## 12. Output & Reporting
- [ ] Generate final markdown/CSV/JSON report
- [ ] Include fields: IP, port, ping, download, upload, score, timestamp
- [ ] Save all raw logs for debugging
- [ ] Create a “good.txt” with just IP:port for quick v2Ray config

## 13. Safety & Etiquette
- [ ] Add random delays (50–200 ms) between requests
- [ ] Limit concurrent scans (e.g., 50–100 threads max)
- [ ] Avoid scanning entire /16 range in < 1 hour
- [ ] Use your own IP or a VPN (to avoid home IP ban)
- [ ] Respect `robots.txt` and CF terms of service (personal use only)

## 14. Optional Enhancements
- [ ] Add IPv6 support
- [ ] Implement SNI brute-force (different Host headers)
- [ ] Test with WebSocket upgrade (`Connection: Upgrade`)
- [ ] Save working IPs to a persistent database
- [ ] Schedule weekly re-scans (IP performance changes over time)
- [ ] Add Telegram/email alert when new fast IP found