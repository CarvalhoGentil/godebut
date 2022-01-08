#!/bin/bash
# PROBLEMA ORIGINAL
# ls -l menu* | head -50 | if [[ (awk '{print $7}') -eq 10 ]]; then echo "nicer" ; fi

# SOLUÇÃO
ls -l deploy* | head -50 | awk '($7 == 10)'
