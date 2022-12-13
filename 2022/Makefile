new:
	@echo "Creating new structure for day $(day) ..."

	@mkdir day$(day)
	@echo "package main" > ./day$(day)/common.go
	@cp template ./day$(day)/part1.go | sed -i -e 's/FuncName/Part1/' -e 's/day/$(day)/' ./day$(day)/part1.go
	@cp template ./day$(day)/part2.go | sed -i -e 's/FuncName/Part2/' -e 's/day/$(day)/' ./day$(day)/part2.go

	@echo "Done!"