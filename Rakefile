require 'rake'
require 'semver'

output = 'dynapi.exe'

task :build do
    buildVersion = SemVer.find.to_s
    buildTimestamp = DateTime.now().strftime("%F %T")

    ldBuildVersion = "-X \"main.buildVersion=#{buildVersion}\""
    ldBuildTimestamp = "-X \"main.buildTimestamp=#{buildTimestamp}\""

    Dir.chdir('cmd') do
        sh('go', 'build', '-ldflags', "#{ldBuildVersion} #{ldBuildTimestamp}", '-o', output)
    end
end

task :run do
    Rake::Task["build"].execute

    Dir.chdir('cmd') do
        sh(output)
    end
end

task :test do
    sh("go test -v")
end

task :cover do
    sh("go test -coverprofile=coverage")
    sh("go tool cover -html=coverage")
end