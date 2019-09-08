#!/bin/bash
cd oracles-disasm; make seasons; cd ..
./oracles-randomizer -dungeons -portals -plan plando.txt 'oracles-disasm/seasons.gbc' plando.gbc
