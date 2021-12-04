PREFIX="Day_"
NUM_DAYS=25
for(( i = 1; i <= $NUM_DAYS; i++)); do
  mkdir "${PREFIX}${i}"
  cd "${PREFIX}${i}"
  curl -o README.md https://adventofcode.com/2021/day/${i}
  cd ..
done