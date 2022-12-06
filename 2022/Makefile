new:
	@echo "Creating new structure for day $(day) ..."

	@mkdir day$(day)
	@echo "package main" > ./day$(day)/common.go
	@cp template ./day$(day)/part1.go | sed 's/day/$(day)/' | sed -i 's/FuncName/Part1'./day$(day)/part1.go
	@cp template ./day$(day)/part2.go | sed 's/day/$(day)/' | sed -i 's/FuncName/Part2'./day$(day)/part1.go

	@echo "Done!"