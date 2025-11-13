import cv2
import json
from pathlib import Path
import numpy as np

def main():
    with open(Path(__file__).parent / "nail_connections.json", "r") as f:
        nail_mappings = json.load(f)

    target_path = Path(__file__).parent / "rune.png"
    
    new_dict = {}
    for k, v in nail_mappings.items():
        new_map = {}
        for kk, vv in v.items():
            if vv > 0:
                new_map[int(kk)] = vv
        new_dict[int(k)] = new_map

    num_points = len(new_dict)

    dim = 800
    img = np.zeros((dim, dim, 3), dtype=np.uint8)

    radius = int(dim*0.49)
    img = cv2.circle(img, (dim//2, dim//2), radius, (200, 200, 200), 5)

    # Calculate angular spacing for evenly distributed nails
    angle_step = 2 * np.pi / num_points

    center_x, center_y = dim // 2, dim // 2
    

    # Draw red markers for each nail position
    nail_positions = {}
    for i in range(num_points):
        angle = (i * angle_step) - (np.pi / 2)
        x = int(center_x + radius * np.cos(angle))
        y = int(center_y + radius * np.sin(angle))
        nail_positions[i+1] = (x, y)

    for nail, connections in new_dict.items():
        for connection in connections.keys():
            start_pos = nail_positions[nail]
            end_pos = nail_positions[connection]
            cv2.line(img, start_pos, end_pos, (0, 255, 0), 1)

    cv2.imwrite(str(target_path), img)


if __name__ == "__main__":
    main()