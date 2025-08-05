from datetime import datetime
from typing import List, Dict
from collections import defaultdict, deque

logs = [
    {"timestamp": "2025-08-05T10:00:01", "ip": "192.168.1.10", "username": "alice", "success": False},
    {"timestamp": "2025-08-05T10:00:03", "ip": "192.168.1.10", "username": "alice", "success": False},
    {"timestamp": "2025-08-05T10:00:06", "ip": "192.168.1.10", "username": "alice", "success": False},
    {"timestamp": "2025-08-05T10:00:09", "ip": "192.168.1.10", "username": "alice", "success": False},
    {"timestamp": "2025-08-05T10:00:15", "ip": "192.168.1.10", "username": "alice", "success": False},
    {"timestamp": "2025-08-05T10:00:25", "ip": "192.168.1.10", "username": "alice", "success": False},
]

def detect_brute_force(logs: List[Dict], threshold: int = 5, window_seconds: int = 60) -> List[str]:
    failed_attempts = defaultdict(list)

    for log in logs:
        if not log["success"]:
            timestamp = datetime.fromisoformat(log["timestamp"])
            failed_attempts[log["ip"]].append(timestamp)

    suspicious_ips = []

    for ip, times in failed_attempts.items():
        times.sort()
        window = deque()
        
        for time in times:
            while window and (time - window[0]).total_seconds() > window_seconds:
                window.popleft()
            window.append(time)
            if len(window) > threshold:
                suspicious_ips.append(ip)
                break

    return suspicious_ips

print(detect_brute_force(logs))