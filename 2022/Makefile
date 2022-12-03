new:
	@echo "Creating new structure for day $(day) ..."

	@mkdir day$(day)
	@cp template ./day$(day)/part1.go | sed -i 's/day/$(day)/' ./day$(day)/part1.go
	@cp template ./day$(day)/part2.go | sed -i 's/day/$(day)/' ./day$(day)/part1.go

	@echo "Done!"